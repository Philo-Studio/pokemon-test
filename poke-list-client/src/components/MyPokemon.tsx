import api from '@/server';
import MyPokemonCard from './MyPokemonCard';

const MyPokemon = async () => {
  const pokemon = await api.getMyPokemon();
  return (
    <>
      <h5>My Pokemon</h5>
      <div className="flex flex-wrap gap-4 bg-gray-500 p-4 rounded-md">
        {pokemon.map((p) => (
          <MyPokemonCard key={p.id} pokemon={p}>
            More content
          </MyPokemonCard>
        ))}
      </div>
    </>
  );
};

export default MyPokemon;
