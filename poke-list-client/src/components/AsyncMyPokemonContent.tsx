import type { MyPokemonItem } from '@/server/types';

const AsyncMyPokemonContent = ({ pokemon }: { pokemon: MyPokemonItem }) => {
  return (
    <div className="flex flex-col gap-4">
      <h3 className="text-lg font-bold">{pokemon.nickname}</h3>
      <p>Load data from server</p>
    </div>
  );
};

export default AsyncMyPokemonContent;
