import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import AuthContextProvider from "@/modules/auth_provider";
import WebSocketProvider from "@/modules/ws_provider";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <AuthContextProvider>
        <WebSocketProvider>
          <body
            className={`${geistSans.variable} ${geistMono.variable} antialiased flex flex-col md:flex-row h-full min-h-screen`}
          >
            {children}
          </body>
        </WebSocketProvider>
      </AuthContextProvider>
    </html>
  );
}
