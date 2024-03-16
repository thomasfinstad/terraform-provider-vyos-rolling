
# Checkpoint
<!--
TODO Continue work on broken create when resource is empty
  add test case as "policy access-list 42"
-->
examples/provider/trace.log:

2024-03-16T22:13:29.396Z [WARN]  provider.terraform-provider-vyos: API Error:: tf_rpc=ReadResource vyos-path=[policy, access-list, 2] @module=vyos error="[api error] 'policy access-list 2': 'Configuration under specified path is empty'" response=<nil> tf_provider_addr=github.com/thomasfinstad/vyos tf_req_id=838e72cb-8aa5-219f-9e3c-6943b65b0946 @caller=/workspaces/terraform-provider-vyos/internal/terraform/helpers/crud/read.go:52 tf_resource_type=vyos_policy_access_list timestamp=2024-03-16T22:13:29.395Z


# Misc
```bash
curl -k --location --request POST "https://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op":"showConfig","path": ["firewall","ipv4","name","example1"]}'&
pid1=$!
curl -k --location --request POST "https://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op":"showConfig","path": ["firewall","ipv4","name","example2"]}'&
pid2=$!

echo -n "Result: "

wait $pid1
ret1=$?
wait $pid2
ret2=$?

echo

exit $(($ret1 + $ret2))
```

```bash
curl -k --location --request POST "https://$VYOS_HOST/configure" --form key="$VYOS_KEY" --form data='[
    {"op":"delete","path":["firewall","ipv4","name","example1","default-action","accept"]},
    {"op":"set","path":["firewall","ipv4","name","example1","default-action","accept"]}
]'&
pid1=$!

curl -k --location --request POST "https://$VYOS_HOST/configure" --form key="$VYOS_KEY" --form data='[
    {"op":"delete","path":["firewall","ipv4","name","example2","default-action","accept"]},
    {"op":"set","path":["firewall","ipv4","name","example2","default-action","accept"]}
]'&
pid2=$!


echo -n "Result: "

wait $pid1
ret1=$?
wait $pid2
ret2=$?

echo

exit $(($ret1 + $ret2))
```

curl -k --location --request POST "https://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op":"showConfig","path":["policy","access-list","2"]}'
