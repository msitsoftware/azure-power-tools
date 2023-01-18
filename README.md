# azure-power-tools (azpt)
Azure Power Tools is a collection of useful developer CLI tools to boost productivity.

# Send Message to Service Bus from CLI

Example:
```bash
azure-power-tools sb send --connectionstring [your_sb_namespace_connection_string] --queue [your_name_of_queue] --contenttype [content_type_of_message_to_send] --message '{\"key\":\"value\"}' 
```

