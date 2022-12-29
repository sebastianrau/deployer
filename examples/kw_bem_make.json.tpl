{
    "steps": [
         {
            "type": "gitUpdate",
            "description": "Updating BasicScripts Repo",
            "ignoreError": true,
            "parameter": {
                "url": 		        "git@github.com:SebastianRau/KW_BEM_BasicScripts.git",
                "dir":              "./temp/KW_BEM_BasicScripts",
                "keyfile":          "./test/keys/id_BasicScripts",
                "password":         "",					
                "ErrorOnNoUpdate": false
            }
        },
        {
            "type": "gitUpdate",
            "description": "Updating Basic Missions Repo",
            "ignoreError": true,
            "parameter": {
                "url": 		        "git@github.com:SebastianRau/KW_BEM_missions.git",
                "dir":              "./temp/KW_BEM_missions",
                "keyfile":          "./test/keys/id_BasicMissions",
                "password":         "",					
                "ErrorOnNoUpdate": false
            }
        },

        {
            "description": "Creating Map Folder",
            "type"       : "mkdir",
            "ignoreError": false,
            "parameter" : {
                "destination": [
                    {{range $i, $map := .mapList}}"./BasicMissions/KW_BEM.{{ $map.MapEnding }}/kw/" {{if hasNext $.mapList $i}},{{end}}
                    {{end}}
                ]
            }
        },


{{range $i, $map := .mapList}}
            {
                "description": "Copy Files for {{ $map.MapName }}",
                "type"       : "cp",
                "ignoreError": false,
                "parameter" : {
                    "src_dest": {
                        "./temp/KW_BEM_missions/KW_BEM.{{ $map.MapEnding }}/*"     : "./BasicMissions/KW_BEM.{{ $map.MapEnding }}/",
                        "./temp/KW_BEM_Missions/BasicFiles/*"                      : "./BasicMissions/KW_BEM.{{ $map.MapEnding }}/",
                        "./temp/KW_BEM_BasicScripts/*"                             : "./BasicMissions/KW_BEM.{{ $map.MapEnding }}/kw/"
                    },
                    "ignoreDotFiles": true
                }
            },    
            {
                "description": "Customize description.ext for {{ $map.MapEnding }}",
                "type"       : "replaceText",
                "ignoreError": false,
                "parameter" : {
                    "destination": "./BasicMissions/KW_BEM.{{ $map.MapEnding }}/description.ext",
                    "map"        : {
                        "%%%briefingName%%%"              	: "KW Basic Event Map {{ $map.MapEnding }}",
                        "%%%onLoadName%%%"                	: "KW Basic Event Map {{ $map.MapEnding }}",
                        "%%%onLoadMission%%%"             	: "[CO44] Kommando Wildsau Basic Event Map {{ $map.MapEnding }}",
                        
                        "%%%loadout_type%%%"              	: "{{ $map.Loadout }}",
                        "%%%loadout_start_equipped%%%"    	: "{{ $map.StartEquipped }}",
                        "%%%loadout_whitelist_enabled%%%" 	: "{{ $map.WhitelistEnabled }}",
                        "%%%civilian_type%%%" 			  	: "{{ $map.CivType }}",
                        "%%%civilian_weapons%%%"            : "{{ $map.CivWeapons }}",
                        "%%%civilian_weapons_chance%%%"     : "{{ $map.CivWeaponsChace }}"
                    }
                }
            }{{if hasNext $.mapList $i}},{{end}}
{{end}}
        ]
    }
