import axios from "axios";

export interface User {
  id: number
username: string
createdAt: string
}

export const axiosReq = axios.create({
  validateStatus: (status: number) => {
    return status >= 200 && status < 300; // default (200 - 299);
  },
  headers: {
    Accept: `application/json`,
    'Content-Type': 'application/json',
  },
  withCredentials: true,
});