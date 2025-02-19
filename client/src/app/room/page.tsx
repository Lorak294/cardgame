import RoomBody from "@/components/room_body";
import { MessageObject } from "@/types/room";
import React, { useState, useRef, useContext, useEffect } from "react";
import { WebSocketContext } from "@/modules/ws_provider";
import { AuthContext } from "@/modules/auth_provider";
import { useRouter } from "next/navigation";
import { API_URL } from "@/constants";
import autosize from "autosize";

const RoomPage = () => {
  // TODO: merge into one room state
  const [messages, setMessages] = useState<Array<MessageObject>>([]);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);

  const textarea = useRef<HTMLTextAreaElement>(null);
  const { connection } = useContext(WebSocketContext);
  const router = useRouter();
  const { user } = useContext(AuthContext);

  // get room info
  useEffect(() => {
    // check for connection
    if (connection == null) {
      router.push("/");
      return;
    }
    const roomId = connection.url.split("/")[5]; // <- TODO: change later
    // call getClients endpoint -> TODO: pull this out of this useeffect handler (if possible because useeffect handler cannot be async)
    const getUsers = async () => {
      try {
        const response = await fetch(`${API_URL}/ws/getClients/${roomId}`, {
          method: "GET",
          headers: { "Content-Type": "application/json" },
        });
        const data = await response.json();
        //console.log("data: " + JSON.stringify(data));
        setUsers(data);
      } catch (error) {
        console.log(error);
      }
    };
    getUsers();
  }, []);

  // handle WS connection communication
  useEffect(() => {
    // attach autosize to textarea
    if (textarea.current) {
      autosize(textarea.current);
    }

    // check for connection
    if (connection == null) {
      router.push("/");
      return;
    }

    // TODO: change the event distinction from string matching to enum or values
    connection.onmessage = (message) => {
      const msg_obj: MessageObject = JSON.parse(message.data);
      if (msg_obj.content == "A new user has joined the room") {
        setUsers([...users, { username: msg_obj.username }]);
      }

      if (msg_obj.content == "User left the room") {
        // delete the user
        const updatedUsers = users.filter(
          (u) => u.username != msg_obj.username
        );
        setUsers(updatedUsers);

        // display the message
        setMessages([...messages, msg_obj]);
        return;
      }

      user?.username == msg_obj.username
        ? (msg_obj.type = "self")
        : (msg_obj.type = "recv");
      setMessages([...messages, msg_obj]);
    };

    // TODO: implement room leaving/joining mechanism
    connection.onclose = () => {};
    connection.onerror = () => {};
    connection.onopen = () => {};
  }, [textarea, messages, connection, users]);

  const sendMessage = (e: React.SyntheticEvent) => {
    if (!textarea.current?.value) return;
    // check for connection
    if (connection === null) {
      router.push("/");
      return;
    }

    connection.send(textarea.current.value);
    textarea.current.value = "";
  };

  return (
    <>
      <div className="flex flex-col w-full">
        <div className="p-4 md:mx-6 mb-14">
          <RoomBody data={messages} />
        </div>
        <div className="fixed bottom-0 mt-4 w-full">
          <div className="flex md:flex-row px-4 py-2 bg-grey md:mx-4 rounded-md">
            <div className="flex w-full mr-4 rounded-md border border-blue">
              <textarea
                ref={textarea}
                placeholder="type your message here"
                className="w-full h-10 p-2 rounded-md focus:outline-none"
                style={{ resize: "none" }}
              />
            </div>
            <div className="flex items-center">
              <button
                className="p-2 rounded-md bg-blue text-white"
                onClick={sendMessage}
              >
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default RoomPage;
