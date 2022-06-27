import { User } from "@/src/types/user";
import { serverURI } from "@/src/api/url";

export const fetchPerpetrator = async (
  perpetrator_id: string
): Promise<User | null> => {
  try {
    const res = await fetch(serverURI + "/users/perpetrator/" + perpetrator_id);

    if (!res.ok) {
      const text = await res.text();
      console.log(text);
      return null;
    }

    const data = await res.json();

    return {
      id: perpetrator_id,
      f_name: data.f_name,
      l_name: data.l_name,
      email: data.email,
      company: data.company,
      position: data.position,
      profile_pic: data.profile_pic,
    };
  } catch (err) {
    console.log(err);
    return null;
  }
};
