{
    "steps": [
        {
            "description": "Write secret file",
            "type"       : "fileWriter",
            "ignoreError": false,
            "parameter" : {
                "filename": "examples/secret.txt",
                "text":[
                    "This is my secret text",
                    "{{secret $.hello $.secretFile}}"
                ]
            }
        }
    ]
}
