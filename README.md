# IKEA_Inventory_Application
its small pilot application to handle ikea inventories

1. clone git repo 
2. docker compose up (it will spin up two instances, one is mysql to store all data, second is ikea app)

3. http://localhost:3000/ikea its welcome page
    

Inventory JSON:
{
  "inventory": [
    {
      "art_id": "1",
      "name": "leg",
      "stock": "12"
    },
    {
      "art_id": "2",
      "name": "screw",
      "stock": "17"
    },
    {
      "art_id": "3",
      "name": "seat",
      "stock": "2"
    },
    {
      "art_id": "4",
      "name": "table top",
      "stock": "1"
    },
    {
      "art_id": "5",
      "name": "table cover",
      "stock": "3"
    }
  ]
}

Product JSON:
{
  "products": [
    {
      "name": "Dining Chair",
      "contain_articles": [
        {
          "art_id": "1",
          "amount_of": "4"
        },
        {
          "art_id": "2",
          "amount_of": "8"
        },
        {
          "art_id": "3",
          "amount_of": "1"
        }
      ]
    },
    {
      "name": "Dining Table",
      "contain_articles": [
        {
          "art_id": "1",
          "amount_of": "4"
        },
        {
          "art_id": "2",
          "amount_of": "8"
        },
        {
          "art_id": "4",
          "amount_of": "1"
        }
      ]
    }
  ]
}