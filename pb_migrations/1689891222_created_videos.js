migrate((db) => {
  const collection = new Collection({
    "id": "mg6r10ygqy6c8m0",
    "created": "2023-07-20 22:13:42.972Z",
    "updated": "2023-07-20 22:13:42.972Z",
    "name": "videos",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "mlwoit1z",
        "name": "title",
        "type": "text",
        "required": true,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
        "system": false,
        "id": "ppisrzcm",
        "name": "source",
        "type": "file",
        "required": true,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "maxSize": 5242880,
          "mimeTypes": [],
          "thumbs": [],
          "protected": false
        }
      }
    ],
    "indexes": [],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("mg6r10ygqy6c8m0");

  return dao.deleteCollection(collection);
})
