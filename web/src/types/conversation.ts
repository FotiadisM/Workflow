export interface Conversation {
  conn_id: string;
  user_id: string;
}

export interface Message {
  id: string;
  senter_id: string;
  text: string;
  time: string;
}
