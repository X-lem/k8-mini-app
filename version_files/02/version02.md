# Version 02

## Routes

| route         | http method | Request                 | Response                                  | Other info                                  |
| ------------- | ----------- | ----------------------- | ----------------------------------------- | ------------------------------------------- |
| /             | GET         |                         | _simple hello response_                   |                                             |
| /pod          | GET         |                         | _name of the k8 pod_                      |                                             |
| /secrets      | GET         |                         | _the data in the secrets yaml file_       | Prints out all the env variables to the log |
| /create-table | POST        |                         | _success message or error will be logged_ | Creates the `users` DB table                |
| /user         | POST        | { "username": "X-lem" } | _returns the created user_                | Creates a user with the given username      |
| /users        | GET         |                         | _returns all the created users_           |                                             |

## Explanation

Expending on the basic setup by adding a DB and a few db call routes.

1. Take a look through the `postgres.yaml` file. You may need to make some small adjustments depending upon your setup
2. I would recommend just deploying the `postgres.yaml` to start so you can see how that's working.

You can see just the DB deployment in action by just deploying the `postgres.yaml` file. Then to connect to it you can use the following command this will connect you to the postgres DB that's on the specified pod (remember you can see the pod name by type `kubectl get po`). All the other information you can get from the yaml file

```
kubectl exec -it [pod-name] --  psql -h localhost -U [user] --password -p [port] [database-name]
```

You will be prompted for the password (because we've added `--password`). If successful you'll get the following.

```
k8-mini-app=#
```

You can start to enter commands (the text below the hyphens is the output you should receive).

```
CREATE TABLE IF NOT EXISTS Users
(
  ID SERIAL NOT NULL PRIMARY KEY,
  Username VARCHAR NOT NULL,
  DateCreated TIMESTAMP NOT NULL DEFAULT NOW()
);
---------
CREATE TABLE
```

```
INSERT INTO Users (Username)
VALUES ('X-lem');
---------
INSERT 0 1
```

```
SELECT * FROM Users;
---------
 id | username |        datecreated
----+----------+----------------------------
  1 | X-lem    | 2023-11-13 17:16:26.871476
(1 row)
```

When you're done you can type `exit` to exit out of the pod. Next feel free to deploy the rest of the `server_*.yaml` files. You should now be able to use the new DB routes (will need to use Postman or something to make the POST requests)

## Suggestions

_Easy_

- Update the golang project to automatically create the `users` table upon start up rather than having to call a route.
- Add a PATCH route to update a user
- Add a DELETE route to delete a user

_Difficult_

- Instead of using the `postgres.yaml` to deploy a DB, create a `Cloud SQL` in gCloud and update the `server_deployment.yaml` to reference it instead.
<details>
<summary>Hints</summary>
  
  - Make sure the Cloud SQL uses a public IP
  - Add `0.0.0.0/0` as an authorized network to the Cloud SQL. This will allow any computer (including the pods) to connect to the DB.
</details>
