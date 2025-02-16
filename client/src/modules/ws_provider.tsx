"use client";
import React, { createContext, useState } from "react";

type Connection = WebSocket | null;

export const WebSocketContext = createContext<{
  connection: Connection;
  setConnection: (c: Connection) => void;
}>({
  connection: null,
  setConnection: () => {},
});

const WebSocketProvider = ({ children }: { children: React.ReactNode }) => {
  const [coonnection, setConnection] = useState<Connection>(null);

  return (
    <WebSocketContext.Provider
      value={{
        connection: coonnection,
        setConnection: setConnection,
      }}
    >
      {children}
    </WebSocketContext.Provider>
  );
};

export default WebSocketProvider;
