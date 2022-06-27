import Navbar from "@/components/navbar/Navbar";
import { Post } from "@/components/homepage/Post";
import { CreatePostArea } from "@/components/homepage/CreatePostArea";
import { AuthRoute, useAuth } from "@/components/auth/AuthRoute";
import { useEffect, useState } from "react";
import { Feed } from "@/src/types/posts";
import { serverURI } from "@/src/api/url";

export default function Home() {
  const auth = useAuth();
  const [feed, setFeed] = useState<Feed[]>([]);

  useEffect(() => {
    (async () => {
      if (auth !== null) {
        if (auth.user !== null) {
          try {
            const res = await fetch(serverURI + "/posts/feed/" + auth.user.id);
            if (!res.ok) {
              const text = await res.text();
              console.log("failed to fetch feed:", text);
            }

            const data = await res.json();
            console.log(data);
            setFeed(data.feed);
          } catch (err) {
            console.log(err);
          }
        }
      }
    })();
  }, []);

  return (
    <AuthRoute>
      <Navbar />
      <main className="relative">
        <div className="absolute transform -translate-x-1/2 left-1/2 mt-10">
          <div className="space-y-4 pb-14">
            <CreatePostArea setFeed={setFeed} />
            {feed === undefined
              ? null
              : [feed.map((f) => <Post key={f.id} feed={f} />)]}
          </div>
        </div>
      </main>
    </AuthRoute>
  );
}
