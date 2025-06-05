import type { AxiosInstance } from 'axios';
import axios from 'axios';
import 'server-only';

class APIServer {
  readonly baseURL: string;
  protected readonly api: AxiosInstance;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
    this.api = axios.create({
      baseURL,
      headers: { 'Content-Type': 'application/json' },
    });
  }

  async getMyPokemon() {
    const response = await this.api.get<
      { id: string; pokedex_id: string; nickname: string; created_at: string; deleted_at: string | null }[]
    >(`/pokemon/mine`);
    return response.data;
  }

  async getAllPokemon() {
    const response = await this.api.get<{ pokedex_id: number; name: string; url: string }[]>(`/pokemon/all`);
    return response.data;
  }
}

if (!process.env.NEXT_PUBLIC_SERVER_URI) {
  throw new Error('NEXT_PUBLIC_SERVER_URI is not set');
}

const api = new APIServer(process.env.NEXT_PUBLIC_SERVER_URI);

export default api;
