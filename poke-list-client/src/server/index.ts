import 'server-only';

import type { AxiosInstance } from 'axios';
import axios from 'axios';
import type { AllPokemonItem, MyPokemonItem } from './types';

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
    const response = await this.api.get<MyPokemonItem[]>(`/pokemon/mine`);
    return response.data;
  }

  async getAllPokemon() {
    const response = await this.api.get<AllPokemonItem[]>(`/pokemon/all`);
    return response.data;
  }
}

if (!process.env.NEXT_PUBLIC_SERVER_URI) {
  throw new Error('NEXT_PUBLIC_SERVER_URI is not set');
}

const api = new APIServer(process.env.NEXT_PUBLIC_SERVER_URI);

export default api;
