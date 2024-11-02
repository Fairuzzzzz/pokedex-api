package poke

var response = `{
    "id": 1,
    "name": "bulbasaur",
    "base_experience": 64,
    "height": 7,
    "weight": 69,
    "order": 1,
    "abilities": [
        {
            "ability": {
                "name": "overgrow"
            },
            "is_hidden": false,
            "slot": 1
        },
        {
            "ability": {
                "name": "chlorophyll"
            },
            "is_hidden": true,
            "slot": 3
        }
    ],
    "moves": [
        {
            "move": {
                "name": "razor-wind"
            }
        },
        {
            "move": {
                "name": "swords-dance"
            }
        }
    ],
    "species": {
        "name": "bulbasaur"
    },
    "stats": [
        {
            "base_stat": 45,
            "effort": 0,
            "stat": {
                "name": "hp"
            }
        },
        {
            "base_stat": 49,
            "effort": 0,
            "stat": {
                "name": "attack"
            }
        },
        {
            "base_stat": 49,
            "effort": 0,
            "stat": {
                "name": "defense"
            }
        },
        {
            "base_stat": 65,
            "effort": 1,
            "stat": {
                "name": "special-attack"
            }
        },
        {
            "base_stat": 65,
            "effort": 0,
            "stat": {
                "name": "special-defense"
            }
        },
        {
            "base_stat": 45,
            "effort": 0,
            "stat": {
                "name": "speed"
            }
        }
    ],
    "types": [
        {
            "slot": 1,
            "type": {
                "name": "grass"
            }
        },
        {
            "slot": 2,
            "type": {
                "name": "poison"
            }
        }
    ]
}`
