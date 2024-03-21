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


Mar 20 20:33:22 vyos02 vyos-http-api[68037]: processing form data
Mar 20 20:33:22 vyos02 vyos-http-api[68037]: processing form data
Mar 20 20:33:22 vyos02 vyos-http-api[68037]: processing form data
Mar 20 20:33:23 vyos02 vyos-http-api[68037]: INFO:      - "POST /retrieve HTTP/1.0" 200 OK
Mar 20 20:33:24 vyos02 vyos-http-api[68037]: INFO:      - "POST /retrieve HTTP/1.0" 200 OK
Mar 20 20:33:25 vyos02 vyos-http-api[68037]: INFO:      - "POST /retrieve HTTP/1.0" 400 Bad Request
Mar 20 20:33:25 vyos02 vyos-http-api[68037]: processing form data
Mar 20 20:33:26 vyos02 vyos-http-api[68037]: INFO:      - "POST /retrieve HTTP/1.0" 200 OK
Mar 20 20:33:26 vyos02 vyos-http-api[68037]: processing form data
Mar 20 20:33:27 vyos02 vyos-http-api[68037]: INFO:      - "POST /retrieve HTTP/1.0" 400 Bad Request
Mar 20 20:33:27 vyos02 vyos-http-api[68037]: processing form data
Mar 20 20:33:27 vyos02 vyos-http-api[68037]: processing form data
Mar 20 20:33:28 vyos02 kernel: vyos-http-api-s[68054]: segfault at 0 ip 00007f9ca3a8b153 sp 00007f9ca34437a0 error 4 in libvyosconfig.so.0[7f9ca399d000+10e000] likely on CPU 0 (core 0, socket 0)
Mar 20 20:33:28 vyos02 systemd[1]: vyos-http-api.service: Main process exited, code=killed, status=11/SEGV
Mar 20 20:33:28 vyos02 systemd[1]: vyos-http-api.service: Failed with result 'signal'.
Mar 20 20:33:28 vyos02 systemd[1]: vyos-http-api.service: Consumed 6.743s CPU time.
Mar 20 20:33:28 vyos02 systemd[1]: vyos-http-api.service: Scheduled restart job, restart counter is at 3.
Mar 20 20:33:28 vyos02 systemd[1]: Stopped vyos-http-api.service - VyOS HTTP API service.
Mar 20 20:33:28 vyos02 systemd[1]: vyos-http-api.service: Consumed 6.743s CPU time.
Mar 20 20:33:28 vyos02 systemd[1]: Started vyos-http-api.service - VyOS HTTP API service.
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: /usr/libexec/vyos/services/vyos-http-api-server:104: PydanticDeprecatedSince20: Pydantic V1 style `@validator` validators are deprecated. You should migrate to Pydantic V2 style `@field_validator` validators, see the migration guide for more details. Deprecated in Pydantic V2.0 to be removed in V3.0. See Pydantic V2 Migration Guide at https://errors.pydantic.dev/2.6/migration/
Mar 20 20:33:29 vyos02 vyos-http-api[68081]:   @validator("path")
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: /usr/share/vyos-http-api-tools/lib/python3.11/site-packages/pydantic/_internal/_config.py:322: UserWarning: Valid config keys have changed in V2:
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: * 'schema_extra' has been renamed to 'json_schema_extra'
Mar 20 20:33:29 vyos02 vyos-http-api[68081]:   warnings.warn(message, UserWarning)
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: Enter main loop...
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: INFO:     Started server process [68081]
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: INFO:     Waiting for application startup.
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: INFO:     Application startup complete.
Mar 20 20:33:29 vyos02 vyos-http-api[68081]: INFO:     Uvicorn running on unix socket /run/api.sock (Press CTRL+C to quit)
