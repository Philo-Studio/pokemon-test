import api from '@/server';
import dynamic from 'next/dynamic';
import MyPokemonCard from './MyPokemonCard';

const DynamicAsyncMyPokemonContent = dynamic(() => import('./AsyncMyPokemonContent'), {
  loading: () => <div>Loading...</div>,
});

const MyPokemon = async () => {
  const pokemon = await api.getMyPokemon();
  return (
    <>
      <h5>My Pokemon</h5>
      <div className="flex flex-wrap gap-4 bg-gray-500 p-4 rounded-md">
        {pokemon.map((p) => (
          <MyPokemonCard key={p.id} pokemon={p}>
            <DynamicAsyncMyPokemonContent pokemon={p} />
          </MyPokemonCard>
        ))}
      </div>
    </>
  );
};

export default MyPokemon;
