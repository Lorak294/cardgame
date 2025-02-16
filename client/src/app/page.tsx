"use client";

import { API_URL } from "@/constants";
import React, { useState, useEffect } from "react";
import { v4 as uuidv4 } from "uuid";

const Home = () => {
  const [roooms, setRooms] = useState<
    { id: string; name: string; playerCount: number }[]
  >([]);
  const [newRoomName, setNewRoomName] = useState("");

  useEffect(() => {
    getRooms();
  }, []);

  const getRooms = async () => {
    try {
      const response = await fetch(`${API_URL}/ws/getRooms`, {
        method: "GET",
      });

      if (response.ok) {
        const data = await response.json();
        setRooms(data);
      }
    } catch (error) {
      console.log(error);
    }
  };

  const createNewRoomHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();

    try {
      setNewRoomName("");
      const response = await fetch(`${API_URL}/ws/createRoom`, {
        method: "POST",
        headers: { "Content-type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          id: uuidv4(),
          name: newRoomName,
        }),
      });

      if (response.ok) {
        getRooms();
      }
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <>
      <div className="my-8 px-4 md:mx-32 w-full h-full">
        <div className="flex justify-center mt-3 p-5">
          <input
            type="text"
            placeholder="room name"
            className="border border-grey p-2 rounded-md focus:outline-none focus:border-blue"
            value={newRoomName}
            onChange={(e) => setNewRoomName(e.target.value)}
          />
          <button
            className="bg-blue border text-white rounded-md p-2 md:ml-4"
            onClick={createNewRoomHandler}
          >
            Create room
          </button>
        </div>
        <div className="mt-6">
          <div className="font-bold"> Avaliable Rooms:</div>
          <div className="grid grid-cols-1 md:grid-cols-5 gap-4 mt-6">
            {roooms.map((room, idx) => (
              <div
                key={idx}
                className="border border-blue p-4 flex items-center rounded-md w-full"
              >
                <div className="w-full">
                  <div className="text-sm">room ({room.playerCount}/10)</div>
                  <div className="text-blue font-bold text-lg">{room.name}</div>
                </div>
                <div className="">
                  <button className="px-4 text-white bg-blue rounded-md">
                    Join
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default Home;
