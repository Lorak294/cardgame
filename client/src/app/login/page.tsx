import React, { useState } from "react";
import { API_URL } from "@/constants";
import { useRouter } from "next/router";

const page = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const router = useRouter();

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();

    try {
      const response = await fetch(`${API_URL}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();
      if (response.ok) {
        const user = {
          username: data.username,
          id: data.id,
        };
        // TODO : change later for something more secure
        localStorage.setItem("user_info", JSON.stringify(user));

        return router.push("/");
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="flex items-center justify-center min-w-full min-h-screen">
      <form className="flex flex-col md:w-1/5">
        <div className="text-3xl font-bold text-center">
          <span className="text-blue">Welcome !</span>
        </div>
        <input
          placeholder="email"
          className="p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <input
          type="password"
          placeholder="password"
          className="p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <button
          className="p-3 mt-6 rounded-md bg-blue font-bold text-whitte"
          type="submit"
          onClick={submitHandler}
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default page;
