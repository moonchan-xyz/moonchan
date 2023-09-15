import { Soup } from "@/functions/soup"

const API_PATH = '/api/'

interface IReply {
  id: number;
  title: string;
  username: string;
  content: string;
}

async function fetchReply(id: number) {
  console.log(id)
  return {
    id : id,
    title: "title",
    username: "username",
    content: "content" + id,
  } as IReply
}

const replySoup = new Soup<number, IReply>(fetchReply)

export type { IReply }
export { API_PATH, fetchReply, replySoup }