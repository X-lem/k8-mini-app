import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { User, axiosReq } from "../api";

const Home: React.FC = () => {
  const [users, setUsers] = useState<User[] | null>(null);

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = () => {
    console.log("fetching users");
    axiosReq
      .get("http://localhost:8080/users")
      .then((res) => {
        console.log("res", res);
        setUsers(res.data);
      })
      .catch((err) => {
        console.log("err", err);
        alert(err.response.data);
      });
  };

  const createUsersTable = () => {
    console.log("creating user table");
    axiosReq
      .post("http://localhost:8080/create-table")
      .then(() => alert("table created!"))
      .catch((err) => {
        console.log("err", err);
        alert(err.response.data);
      });
  };

  return (
    <main>
      <Link to="/create-user">Create User</Link>
      <h1>Home Page</h1>
      <button onClick={createUsersTable}>Create Users Table</button>
      {/* <ul>
        {users.map((user, index) => (
          <li key={index}>{user}</li>
        ))}
      </ul> */}
    </main>
  );
};

export default Home;
