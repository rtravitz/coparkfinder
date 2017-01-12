This repository is part of a larger project that also includes:
* [Frontend](https://github.com/rtravitz/coparkfinder-front)
* [Alexa Skill](https://github.com/rtravitz/coparkfinder-alexa)

The API is available in production [on Heroku](https://coparkfinder.herokuapp.com/).

-------------------
## Overview
Colorado Parks and Wildlife publishes a large amount of information about the 42 parks it controls. While the information is available, it is difficult to find exactly what you are looking for based on your interests and the facilities available. This Go API provides the information in an easily accessible JSON format.

## Endpoints
### /api/v1/activities
This endpoint returns an array of activities in the following format:
```
[
  {
    "id": 1,
    "type": "biking"
  },
  {
    "id": 2,
    "type": "boating"
  },
  {
    "id": 23,
    "type": "boating (nonmotorized)"
  }
]
```

### /api/v1/facilities
This endpoint returns an array of facilities in the following format:
```
[
  {
    "id": 8,
    "type": "archery range"
  },
  {
    "id": 19,
    "type": "archery/shooting range"
  },
  {
    "id": 3,
    "type": "basic campsites"
  }
]
```

### /api/v1/parks
If no query parameters are sent, this endpoint will return a list of all parks. The endpoint supports queries for ```activties``` and ```facilities```. Lists of parameters must be comma separated and have single quotes (%27) surrounding each item.

Example queries:

* ``` /api/v1/parks?activities='boating' ```
* ``` /api/v1/parks?facilities='visitor%20center','boat%20ramp' ```
* ``` /api/v1/parks?facilities='visitor%center'&activities='boating','fishing' ```


The endpoint will return an array of parks in the following format:

```
[
  {
    "id": 1,
    "name": "Arkansas Headwaters Recreation Area",
    "street": "307 W. Sackett Ave.",
    "city": "Salida",
    "zip": "81201",
    "email": "ahra@state.co.us",
    "description": "\"We are the river.\" That's the catchphrase for fun and adventure in this recreation area along one of the most popular whitewater boating rivers in the United States, etc",
    "url": "http://cpw.state.co.us/placestogo/Parks/ArkansasHeadwatersRecreationArea",
    "facilities": [
                    {
                      "id": 1,
                      "type": "visitor center"
                    },
                    {
                      "id": 2,
                      "type": "boat ramp"
                    }
                  ],
    "activities": [
                     {
                       "id": 1,
                        "type": "biking"
                     },
                     {
                       "id": 2,
                       "type": "boating"
                     }
                  ]
   }
]
```

### /api/v1/park
To access a park at this endpoint, you must query for a park name. 

For example:
* ```/api/v1/park?name=Arkansas%20Headwaters%20Recreation%20Area``` or ```/api/v1/park?name=Crawford```

The endpoint will return a park in the following format:

```
{
  "id": 1,
  "name": "Arkansas Headwaters Recreation Area",
  "street": "307 W. Sackett Ave.",
  "city": "Salida",
  "zip": "81201",
  "email": "ahra@state.co.us",
  "description": "\"We are the river.\" That's the catchphrase for fun and adventure in this recreation area along one of the most popular whitewater boating rivers in the United States, etc",
  "url": "http://cpw.state.co.us/placestogo/Parks/ArkansasHeadwatersRecreationArea",
  "facilities": [
                  {
                    "id": 1,
                    "type": "visitor center"
                  },
                  {
                    "id": 2,
                    "type": "boat ramp"
                  }
                ],
  "activities": [
                   {
                     "id": 1,
                      "type": "biking"
                   },
                   {
                     "id": 2,
                     "type": "boating"
                   }
                ]
 }
```
