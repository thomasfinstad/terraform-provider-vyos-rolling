# TODO large steps
- Verify CRUD functions
- Create custom import function that will populate `id` and `*_id` fields correctly
- Auto generate resource import docs
- Fix doc "subcategory"
- Create global (Node) resources
- improve error vs diag returns (especially (un)marshal helpers)


# Misc
curl -k --location --request POST "https://$VYOS_HOST/configure" --form key="$VYOS_KEY" --form data='[
    {"op":"set","path":["firewall","zone","TF-Examples","intra-zone-filtering","action","accept"]},
    {"op":"set","path":["firewall","zone","TF-Examples","intra-zone-filtering","firewall","name","test"]}
]'

curl -k --location --request POST "https://$VYOS_HOST/retrieve" --form key="$VYOS_KEY" --form data='{"op": "showConfig", "path": ["firewall"]}'
