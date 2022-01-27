# Aceh Dictionary API

 الْحَمْدُ لِلَّهِ الَّذِى بِنِعْمَتِهِ تَتِمُّ الصَّالِحَاتُ 
 
 Segala puji bagi Allah, dengan segala kenikmatannya sempuranalah segala kebaikan.
 
 Project ini pada dasarnya merupakan kamus bahasa aceh. Selain itu, project ini berguna membantu memberi saran kosa kata yang dimaksud kepada pengguna untuk mengetik kosa kata bahasa aceh yang benar. 

 Pada project ini diterapkan algoritma Jaro Winkler untuk menentukan similiarity karakter kata yang ditulis oleh pengguna dengan data di database.


## APIs

### Get Advices

Request 🔥
- Method : POST
- Endpoint : /api/v1/advices
- Header :
- Content_Type : application/json
- Accept : application/json
- Params :
- Body :

```json
{
    "input": "lon"
}
```
Response 🚀
```json
{
  "meta": {
    "message": "Successfully to get word advice",
    "code": 200,
    "status": "success"
  },
  "data": [
    {
      "Aceh": "lo",
      "Indonesia": "muat",
      "Similiarity": 0.9111111111111111
    },
    {
      "Aceh": "lonton",
      "Indonesia": "tonton",
      "Similiarity": 0.8833333333333334
    },
    {
      "Aceh": "loncèng",
      "Indonesia": "lonceng",
      "Similiarity": 0.8666666666666668
    },
    {
      "Aceh": "glong",
      "Indonesia": "tancap",
      "Similiarity": 0.8666666666666667
    },
    {
      "Aceh": "glong",
      "Indonesia": "lantak",
      "Similiarity": 0.8666666666666667
    }
  ]
}
```

## Credit

- Data kamus di ambil dengan teknik scraping dari [kata.web.id](https://kata.web.id/)

## Support Me

- 🚀 [Paypal](https://paypal.me/sahibulnf)


Copyright © 2022 by [Sahibul Nuzul Firdaus](https://sahibul-nf.github.io/)