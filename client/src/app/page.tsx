"use client";

import { API_URL, WEBSOCKET_URL } from "@/constants";
import React, { useState, useEffect, useContext } from "react";
import { v4 as uuidv4 } from "uuid";
import { AuthContext } from "@/modules/auth_provider";
import { WebSocketContext } from "@/modules/ws_provider";
import { useRouter } from "next/navigation";

const Home = () => {
  const [roooms, setRooms] = useState<
    { id: string; name: string; playerCount: number }[]
  >([]);
  const [newRoomName, setNewRoomName] = useState("");
  const { user } = useContext(AuthContext);
  const { setConnection } = useContext(WebSocketContext);
  const router = useRouter();

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

  const joinRoom = (roomId: string) => {
    const ws = new WebSocket(
      `${WEBSOCKET_URL}/ws/joinRoom/${roomId}?userId=${user.id}&username=${user.username}`
    );

    if (ws.OPEN) {
      setConnection(ws);
      router.push("/room");
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
                  <button
                    className="px-4 text-white bg-blue rounded-md"
                    onClick={() => joinRoom(room.id)}
                  >
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
