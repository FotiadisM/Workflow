export interface Feed {
  id: string;
  type: "post" | "share" | "comment" | "like";
  post_id: string;
  perpetator_id: string;
}

export interface Post {
  id: string;
  user_id: string;
  text: string;
  images: string[];
  videos: string[];
  likes: string[];
  comments: string[];
  created: string;
}

export interface Comment {
  id: string;
  user_id: string;
  text: string;
  created: string;
  likes: string[];
}
