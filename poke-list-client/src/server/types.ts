export type MyPokemonItem = {
  id: string;
  pokedex_id: string;
  nickname: string;
  created_at: string;
  deleted_at: string | null;
};

export type AllPokemonItem = { pokedex_id: number; name: string; url: string };
