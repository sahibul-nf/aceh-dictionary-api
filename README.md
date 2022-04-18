
# Acehnese Dictionary

Praise be to Allah, with all the pleasures all goodness is perfect.

This project is basically an aceh language dictionary. In addition, this project is useful in helping to suggest the intended vocabulary for users to type the correct Acehnese vocabulary.

In this project, the Jaro Winkler algorithm is applied to determine the similiarity of the word characters written by the user with the data in the database.
## Features

- [x]   Show list of acehnese dictionary data
- [x]   Word suggestion list in aceh language


## Tech Stack

- Go
- GORM
- Gin Web Framework
- PostgreSQL
- Rest API
- Heroku



## API Reference

### Get all data

```http
  GET /api/v1/dictionaries
```

### Get Advices

Request 🔥
- Endpoint : 
```http
  GET /api/v1/advices?input=peu
```
- Params : `input`
- Body : none
  
Response 🚀
```json
// https://aceh-dictionary.herokuapp.com/api/v1/advices?input=peu

{
  "meta": {
    "message": "Successfully to get word advice",
    "code": 200,
    "status": "success"
  },
  "data": [
    {
      "id": 188,
      "aceh": "peue",
      "indonesia": "apa",
      "similiarity": 0.9416666666666667
    },
    {
      "id": 654,
      "aceh": "peuja",
      "indonesia": "boraks",
      "similiarity": 0.9066666666666667
    },
    {
      "id": 4063,
      "aceh": "peuda",
      "indonesia": "pada",
      "similiarity": 0.9066666666666667
    },
    {
      "id": 1338,
      "aceh": "peuet",
      "indonesia": "empat",
      "similiarity": 0.9066666666666667
    },
    {
      "id": 4432,
      "aceh": "peuta",
      "indonesia": "peta",
      "similiarity": 0.9066666666666667
    }
  ]
}
```
## Data Source

- https://kata.web.id := Aceh - Indonesia
<!-- ## License

[MIT](https://choosealicense.com/licenses/mit/) -->


## Support

For **Global** support, you can buy a coffee.

[![Support](https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png)](https://www.buymeacoffee.com/sahibulnf)

For **Indonesian** support,
- [Karyakarsa](https://karyakarsa.com/sahibul_nf)
