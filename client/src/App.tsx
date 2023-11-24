import "./App.css";
import {
  Navigate,
  Outlet,
  RouterProvider,
  createBrowserRouter,
} from "react-router-dom";
import Home from "./components/Home";
import CreateUser from "./components/CreateUser";

function App() {
  const router = createBrowserRouter([
    {
      element: <Outlet />,
      children: [
        {
          path: "/",
          element: <Home />,
        },
        {
          path: "/create-user",
          element: <CreateUser />,
        },
      ],
      errorElement: <Navigate to={"/"} />,
    },
  ]);

  return (
    <div id="k8mini-app" className="k8mini-app">
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
