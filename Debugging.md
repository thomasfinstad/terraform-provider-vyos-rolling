# Debugging

This is a little journal of one of the rabbit holes I dived into during one of my depressive nights of insomnia. It goes as it often does when you work with stuff you aren't too familiar with, but it was a fun little ride.
I will keep the story here as a bit of entertainment and a reminder to all that shit be whack yo.

## Problem description

When running `go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/` the `TestCrudDeleteRetrySuccess` test times out, but not when it is run alone: `go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ -run TestCrudDeleteRetrySuccess`.

The timeout is a bug, as sleep is called with a 520ms time.Duration struct, but does not stop sleeping and is terminated by the timeout while in the sleep function.

Log snippet:
```
2024/03/20 14:06:04 Field: ExistsTagFirewallIPvfourNameRule     Type: bool      Tags: [rule child]
2024/03/20 14:06:04 Total retry delay: 0s
2024/03/20 14:06:04 Total retry count: 0
2024/03/20 14:06:04 Retry delay..: 520ms
panic: test timed out after 5s
running tests:
        TestCrudDeleteRetrySuccess (0s)

goroutine 1442 [running]:
testing.(*M).startAlarm.func1()
        /usr/local/go/src/testing/testing.go:2366 +0x385
created by time.goFunc
        /usr/local/go/src/time/sleep.go:177 +0x2d

goroutine 1 [chan receive]:
testing.(*T).Run(0xc00014e4e0, {0x96252b?, 0x0?}, 0x99ec98)
        /usr/local/go/src/testing/testing.go:1750 +0x3ab
testing.runTests.func1(0xc00014e4e0)
        /usr/local/go/src/testing/testing.go:2161 +0x37
testing.tRunner(0xc00014e4e0, 0xc000175c70)
        /usr/local/go/src/testing/testing.go:1689 +0xfb
testing.runTests(0xc0001240c0, {0xd54100, 0x11, 0x11}, {0x1?, 0x50c2ae?, 0xd5aa80?})
        /usr/local/go/src/testing/testing.go:2159 +0x445
testing.(*M).Run(0xc000121040)
        /usr/local/go/src/testing/testing.go:2027 +0x68b
main.main()
        _testmain.go:79 +0x16c

goroutine 1273 [sleep]:
time.Sleep(0x1efe9200)
        /usr/local/go/src/runtime/time.go:195 +0x115
github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/crud.delete({_, _}, {{{{0xa35340, 0xc000178140}, 0x0, {0x0, 0x0}, 0x0}, {0x959297, 0xa}, ...}, ...}, ...)
        /workspaces/terraform-provider-vyos/internal/terraform/helpers/crud/delete.go:137 +0x665
github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/crud.TestCrudDeleteRetrySuccess(0xc00046a000)
        /workspaces/terraform-provider-vyos/internal/terraform/helpers/crud/delete_test.go:492 +0x79a
testing.tRunner(0xc00046a000, 0x99ec98)
        /usr/local/go/src/testing/testing.go:1689 +0xfb
created by testing.(*T).Run in goroutine 1
        /usr/local/go/src/testing/testing.go:1742 +0x390
```

Related code snippet from `internal/terraform/helpers/crud/delete.go:137` :
```
					log.Println("Total retry delay:", retryTotalDelay)
					log.Println("Total retry count:", retryCnt)
					log.Println("Retry delay..:", backOff)
					tflog.Info(ctx, "delaying before next retry", map[string]interface{}{"retryTotalDelay": retryTotalDelay, "retryCnt": retryCnt, "backOff": backOff})
					time.Sleep(backOff)
					log.Println("Retrying...")
					retryTotalDelay += backOff
					retryCnt++
```

## Debugging steps

## Desperate mode

Since the issue does not show up no matter what when the test is run by itself with or without debug mode it is hard to nail down.

Maybe it is because the http server is a real server, causing issues, but that seems very unlikely.

Next step will be to try to nail down the minimum set of test combos needed to make the issue appear.

list tests, excluding the problematic one:
```
$ go test ./internal/terraform/helpers/crud -list . | grep -v "^ok \|TestCrudDeleteRetrySuccess" > tmp/tests
$ wc -l tmp/tests
16 tmp/tests
```

16 tests, narrowing this down to the problematic combo should be easy!

use py script to create permutations of tests:

```
from os import read
import sys
from itertools import permutations, chain, combinations
from time import sleep

import operator
from collections import Counter
from functools import reduce
from math import factorial


if len(sys.argv) <= 1:
    print('No arguments provided; exiting.', file=sys.stderr)
    sys.exit()

args = sys.argv[1:]
print("args:", args, file=sys.stderr)
count = 0

if args.count("TestCrudDeleteRetrySuccess"):
    args.remove("TestCrudDeleteRetrySuccess")

totalPermutations = 0
for i in range(1, len(args)+1):
    cntPerms=0
    for combo in combinations(args, i):
        comboToLength =(combo)[:i]
        num = factorial(len(comboToLength))
        mults = Counter(comboToLength).values()
        den = reduce(operator.mul, (factorial(v) for v in mults), 1)
        cntPerms +=int(num / den)
    totalPermutations += cntPerms
    print(f"{i} tests has permutations: {cntPerms:,} currently total tallied permutations: {totalPermutations:,}", file=sys.stderr)

print("Press [Enter] if you want to produce all these permutations, [Ctrl-C] to cancel", file=sys.stderr)
input()

sudoSandwich = chain(permutations(args, x) for x in range(1, len(args)+1))

curLen = 0
doneCnt=0
for combo in sudoSandwich:
    for e in combo:
        if curLen != len(e):
            curLen = len(e)
            print("Current permutation length:", curLen, file=sys.stderr)
        doneCnt +=1
        if doneCnt % 1000000 == 0:
            print(f"[{int(doneCnt/totalPermutations)*100}%] {doneCnt}/{totalPermutations}", file=sys.stderr)
        print("|".join(e)+"|TestCrudDeleteRetrySuccess")
```

No I will not take critisim on the script, its is a expertly crafted masterpeice of abstract art!


```
$ time timeout 10s python tmp/perm.py $(cat tmp/tests) > tmp/perms

args: ['TestCrudCreateSuccess', 'TestCrudCreateResourceAlreadyExsitsFailure', 'TestCrudCreateResourceAlreadyExsitsIgnore', 'TestCrudCreateResourceParentMissingFailure', 'TestCrudCreateResourceParentMissingIgnore', 'TestCrudCreateTimeoutSuccess', 'TestCrudCreateTimeoutFailure', 'TestCrudCreateRetrySuccess', 'TestCrudDeleteSuccess', 'TestCrudDeleteResourceHasChildFailure', 'TestCrudDeleteResourceHasChildIgnore', 'TestCrudDeleteGlobalResourceWithChild', 'TestCrudDeleteGlobalResourceWithoutChild', 'TestCrudReadSuccess', 'TestCrudReadEmptyResource', 'TestCrudUpdateCrossResourceContamination']

1 tests has permutations: 16 currently total tallied permutations: 16
2 tests has permutations: 240 currently total tallied permutations: 256
3 tests has permutations: 3,360 currently total tallied permutations: 3,616
4 tests has permutations: 43,680 currently total tallied permutations: 47,296
5 tests has permutations: 524,160 currently total tallied permutations: 571,456
6 tests has permutations: 5,765,760 currently total tallied permutations: 6,337,216
7 tests has permutations: 57,657,600 currently total tallied permutations: 63,994,816
8 tests has permutations: 518,918,400 currently total tallied permutations: 582,913,216
9 tests has permutations: 4,151,347,200 currently total tallied permutations: 4,734,260,416
10 tests has permutations: 29,059,430,400 currently total tallied permutations: 33,793,690,816
11 tests has permutations: 174,356,582,400 currently total tallied permutations: 208,150,273,216
12 tests has permutations: 871,782,912,000 currently total tallied permutations: 1,079,933,185,216
13 tests has permutations: 3,487,131,648,000 currently total tallied permutations: 4,567,064,833,216
14 tests has permutations: 10,461,394,944,000 currently total tallied permutations: 15,028,459,777,216
15 tests has permutations: 20,922,789,888,000 currently total tallied permutations: 35,951,249,665,216
16 tests has permutations: 20,922,789,888,000 currently total tallied permutations: 56,874,039,553,216
Press [Enter] if you want to produce all these permutations, [Ctrl-C] to cancel

Current permutation length: 1
Current permutation length: 2
Current permutation length: 3
Current permutation length: 4
Current permutation length: 5
Current permutation length: 6
[0%] 1000000/56874039553216
[0%] 2000000/56874039553216
[0%] 3000000/56874039553216
[0%] 4000000/56874039553216
[0%] 5000000/56874039553216
[0%] 6000000/56874039553216
Current permutation length: 7
[0%] 7000000/56874039553216
[0%] 8000000/56874039553216
[0%] 9000000/56874039553216
[0%] 10000000/56874039553216

real    0m10.509s
user    0m7.718s
sys     0m1.991s
```

What, 16 isnt that many test, damn math always creating big numbers for nothing.

```
$ wc -l tmp/perms
10233763 tmp/perms

$ head -n 2000 tmp/perms | sort | uniq -d
```

Damn, no duplicates from the first 2k lines, was really hoping I messed up badly.

This script is slow, but how long would it take to complete?

```
$ jq -n 56874039553216/10000000*10/60/60/24/365
1.8034639635088785
```

1.8 years.... okae we must find another way if we want to do this before terraform is depriciated and we all live on mars...

## Traceing

Maybe we can try to find something through tracing?

`go test -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ -trace trace.out`

`go tool trace trace.out`

Hmm, web page? neat... I think...

No... if tracing could help we need someone better at it than me.

## It is made out of meat

Time to turn on the meat computer and not be so damn stupid about which tests to run.

Lets just list the tests and the order they are being run in

`go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ | grep "^=== RUN" | awk '{print $3}' > tmp/problemTests`

```
$ cat tmp/problemTests
TestCrudCreateSuccess
TestCrudCreateResourceAlreadyExsitsFailure
TestCrudCreateResourceAlreadyExsitsIgnore
TestCrudCreateResourceParentMissingFailure
TestCrudCreateResourceParentMissingIgnore
TestCrudCreateTimeoutSuccess
TestCrudCreateTimeoutFailure
TestCrudCreateRetrySuccess
TestCrudDeleteSuccess
TestCrudDeleteResourceHasChildFailure
TestCrudDeleteResourceHasChildIgnore
TestCrudDeleteGlobalResourceWithChild
TestCrudDeleteGlobalResourceWithoutChild
TestCrudDeleteRetrySuccess
```

Lets see how many we can skip

```
$ for i in $(seq 2 $(wc -l <tmp/problemTests)); do echo $i; go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ -run "$(cat tmp/problemTests | tail -n "$i" | tr '\n' '|' | head -c -1)" >/dev/null || break; done
2
3
4
5
6
7
```

So we know this test combo will still error out:

```
$ cat tmp/problemTests | tail -n "7"
TestCrudCreateRetrySuccess
TestCrudDeleteSuccess
TestCrudDeleteResourceHasChildFailure
TestCrudDeleteResourceHasChildIgnore
TestCrudDeleteGlobalResourceWithChild
TestCrudDeleteGlobalResourceWithoutChild
TestCrudDeleteRetrySuccess
```

Lets get the permutations

```
$ python tmp/perm.py $(cat tmp/reduced) > tmp/permsReduced
$ wc -l tmp/permsReduced
1956 tmp/permsReduced
```

We have reduced it to 0.0000000034% of the original stupidity.

```
$ jq -n 1956/56874039553216*100
3.4391789564548276e-09
```

Still 2k permutations means we might never finish. Im out of "clever" ideas, lets just have it run until we find the smallest combo of failing tests, or we die of old age...

```
t=$(wc -l<tmp/permsReduced); i=0; for l in $(cat tmp/permsReduced); do echo "$((i++))/$t $l"; go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ -test.run "$l" >/dev/null || break; done
0/1956 TestCrudCreateRetrySuccess|TestCrudDeleteRetrySuccess
1/1956 TestCrudDeleteSuccess|TestCrudDeleteRetrySuccess
2/1956 TestCrudDeleteResourceHasChildFailure|TestCrudDeleteRetrySuccess
3/1956 TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteRetrySuccess
4/1956 TestCrudDeleteGlobalResourceWithChild|TestCrudDeleteRetrySuccess
5/1956 TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteRetrySuccess
6/1956 TestCrudCreateRetrySuccess|TestCrudDeleteSuccess|TestCrudDeleteRetrySuccess
7/1956 TestCrudCreateRetrySuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteRetrySuccess
```

No way!... 7 runs to find it? Did I do all that for so little? That cant be right...

```
$ time go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ -run "^TestCrudCreateRetrySuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteRetrySuccess$" | grep "^[-=]\|
^panic"
=== RUN   TestCrudCreateRetrySuccess
--- PASS: TestCrudCreateRetrySuccess (1.69s)
=== RUN   TestCrudDeleteResourceHasChildFailure
--- PASS: TestCrudDeleteResourceHasChildFailure (2.26s)
=== RUN   TestCrudDeleteRetrySuccess
panic: test timed out after 5s

real    0m5.325s
user    0m0.550s
sys     0m0.301s
```

Balls, it is, that was a lot of wasted effort. So the issue appears when these three tests runs after one another:

1. TestCrudCreateRetrySuccess
2. TestCrudDeleteResourceHasChildFailure
3. TestCrudDeleteRetrySuccess

## Any light at the end?

So now we know the minimum tests required to generate the issue, I wonder if the issue appears when not running those two tests right before that failing one. Lets give it a small try just to check.

I have step a way for a bit so might as well give it 10 min.

```
$ timeout 10m bash -c 't=$(wc -l<tmp/permsReduced); i=0; for l in $(cat tmp/permsReduced); do [[ "$l" == *"TestCrudCreateRetrySuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteRetrySuccess" ]] && continue; echo "$((i++))/$t $l"; go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ -test.run "$l" >/dev/null || break; done'
0/1956 TestCrudCreateRetrySuccess|TestCrudDeleteRetrySuccess
1/1956 TestCrudDeleteSuccess|TestCrudDeleteRetrySuccess
2/1956 TestCrudDeleteResourceHasChildFailure|TestCrudDeleteRetrySuccess
3/1956 TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteRetrySuccess
4/1956 TestCrudDeleteGlobalResourceWithChild|TestCrudDeleteRetrySuccess
5/1956 TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteRetrySuccess
6/1956 TestCrudCreateRetrySuccess|TestCrudDeleteSuccess|TestCrudDeleteRetrySuccess
7/1956 TestCrudCreateRetrySuccess|TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteRetrySuccess
8/1956 TestCrudCreateRetrySuccess|TestCrudDeleteGlobalResourceWithChild|TestCrudDeleteRetrySuccess
9/1956 TestCrudCreateRetrySuccess|TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteRetrySuccess
10/1956 TestCrudDeleteSuccess|TestCrudCreateRetrySuccess|TestCrudDeleteRetrySuccess
11/1956 TestCrudDeleteSuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteRetrySuccess
12/1956 TestCrudDeleteSuccess|TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteRetrySuccess
13/1956 TestCrudDeleteSuccess|TestCrudDeleteGlobalResourceWithChild|TestCrudDeleteRetrySuccess
14/1956 TestCrudDeleteSuccess|TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteRetrySuccess
15/1956 TestCrudDeleteResourceHasChildFailure|TestCrudCreateRetrySuccess|TestCrudDeleteRetrySuccess
```

That didn't take long. So the ordering of the "prerequsites" does not matter it seems, is there any other combo than these 3 tests that causes the issues?

(truncated)

```
$ timeout 10m bash -c 't=$(wc -l<tmp/permsReduced); i=0; for l in $(cat tmp/permsReduced); do [[ "$l" == *"TestCrudCreateRetrySuccess"* && "$l" == *"TestCrudDeleteResourceHasChildFailure"* ]] && continue; echo "$((i++))/$t $l"; go test -v -count 1 -failfast -timeout 5s  ./internal/terraform/helpers/crud/ -test.run "$l" >/dev/null || break; done'
0/1956 TestCrudCreateRetrySuccess|TestCrudDeleteRetrySuccess
1/1956 TestCrudDeleteSuccess|TestCrudDeleteRetrySuccess
2/1956 TestCrudDeleteResourceHasChildFailure|TestCrudDeleteRetrySuccess
3/1956 TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteRetrySuccess
[...]
161/1956 TestCrudDeleteSuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteRetrySuccess
162/1956 TestCrudDeleteSuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteGlobalResourceWithChild|TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteRetrySuccess
163/1956 TestCrudDeleteSuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteGlobalResourceWithChild|TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteRetrySuccess
164/1956 TestCrudDeleteSuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteResourceHasChildIgnore|TestCrudDeleteRetrySuccess
165/1956 TestCrudDeleteSuccess|TestCrudDeleteResourceHasChildFailure|TestCrudDeleteGlobalResourceWithoutChild|TestCrudDeleteGlobalResourceWithChild|TestCrudDeleteRetrySuccess
```

Alright that is good enough, I think the 2 tests we found must BOTH be before the failing test.

I will try to run the tests in debug mode to see if I can inspect close to the code as it fails.

```
$ mv internal/terraform/helpers/crud/create_test.go internal/terraform/helpers/crud/create_test_a.go
$ mv internal/terraform/helpers/crud/delete_test.go internal/terraform/helpers/crud/delete_test_a.go
$ cat tmp/borked.go
package main

import (
        "testing"

        "github.com/thomasfinstad/terraform-provider-vyos/internal/terraform/helpers/crud"
)

func main() {
        t1 := &testing.T{}
        crud.TestCrudCreateRetrySuccess(t1)
        t2 := &testing.T{}
        crud.TestCrudDeleteResourceHasChildFailure(t2)
        t3 := &testing.T{}
        crud.TestCrudDeleteRetrySuccess(t3)
}
```

Well, while the output is a bitch, at no point does the test just hang forever so what ever is going on seems to be part of a full and balanced breakf... erhm... test run.

What now? I am running out of ideas and steam... At least lets revert the test file renaming.

```
$ mv internal/terraform/helpers/crud/create_test_a.go internal/terraform/helpers/crud/create_test.go
$ mv internal/terraform/helpers/crud/delete_test_a.go internal/terraform/helpers/crud/delete_test.go
```

Yup, back to a working problem when running the tests.

I guess it is time to just look at the tests and see if there is anything I can just intuit to be wrong when the tests are run in this order:

1. Pre-reqs in either order:
    * TestCrudCreateRetrySuccess
    * TestCrudDeleteResourceHasChildFailure
2. TestCrudDeleteRetrySuccess

Did you know that humans are stupid? I did, well at least I know of one that is.

You see I don't know why, but I do know that I assumed that the parameter `-timeout 5s` were a per test setting. And now you know where this is going.
Yeah, it is not. It is a per execution setting. I will set the timeout to be 30 seconds instead. That is plenty for the few tests I have.

Moral of the story: don't waste hours and hours because you skipped reading the manual for one of the parameters you use.
