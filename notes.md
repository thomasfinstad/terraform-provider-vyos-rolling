
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
