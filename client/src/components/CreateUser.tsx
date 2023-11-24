import React, { useState } from "react";
import { Link } from "react-router-dom";
import { axiosReq } from "../api";

const CreateUser: React.FC = () => {
  const [username, setUsername] = useState("");

  const createUser = () => {
    if (username === "") return;
    
    console.log(`Creating user: ${username}`);
    axiosReq
      .post("http://localhost:8080/user", {username: username})
      .then((res) => {
        console.log("res", res);
        alert("User created! " + res.data.toString())
      })
      .catch((err) => {
        console.log("err", err);
        alert(err.response.data);
      });
  };

  return (
    <main>
      <Link to="/">Home</Link>
      <h1>Create User</h1>
      <input
        type="text"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        placeholder="Enter username"
      />
      <button onClick={createUser}>Create</button>
    </main>
  );
};

export default CreateUser;
