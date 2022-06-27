import { serverURI } from "@/src/api/url";
import { User } from "@/src/types/user";
import { useRouter } from "next/router";
import { useState } from "react";

export function useProvideAuth() {
  const router = useRouter();
  const [user, setUser] = useState<User | null>(null);

  const signIn = async (
    email: string,
    password: string
  ): Promise<string | null> => {
    if (email !== "" && password !== "") {
      try {
        const res = await fetch(serverURI + "/auth/signIn", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            email: email,
            password: password,
          }),
        });

        if (res.status === 200) {
          const u = await res.json();
          if (u.user !== undefined) {
            setUser(u.user);
            return null;
          }
        }

        return "error";
      } catch (err) {
        return "error";
      }
    }

    //     setUser({
    //       id: "0e306495-8cc6-443e-94c5-b6585667c8e7",
    //       f_name: "Thodoris",
    //       l_name: "Flebas",
    //       email: "flebas@mail.com",
    //       company: "Skalvenitis",
    //       position: "Tameias",
    //       profile_pic: "2777baee-5177-4ce9-8ffb-726adb4c39cb",
    //       role: "admin",
    //     });

    return null;
  };

  const signUp = async (
    f_name: string,
    l_name: string,
    email: string,
    company: string,
    position: string,
    password: string,
    profile: File
  ): Promise<string | null> => {
    const formData = new FormData();
    formData.append("f_name", f_name);
    formData.append("l_name", l_name);
    formData.append("email", email);
    formData.append("company", company);
    formData.append("position", position);
    formData.append("password", password);
    formData.append("profile", profile, profile.name);

    const res = await fetch(serverURI + "/auth/signUp", {
      method: "POST",
      body: formData,
    });

    if (!res.ok) {
      const text = await res.text();
      if (text === "key already exist") return "Email is already in use";
      return res.statusText + " " + text;
    }

    const data = await res.json();
    setUser({
      id: data.user.id,
      f_name: data.user.f_name,
      l_name: data.user.l_name,
      email: data.user.email,
      company: data.user.company,
      position: data.user.position,
      profile_pic: data.user.profile_pic,
      role: "normal",
    });

    return null;
  };

  const signOut = () => {
    setUser(null);
    router.push("/");
  };

  return {
    user,
    setUser,
    signIn,
    signUp,
    signOut,
  };
}
