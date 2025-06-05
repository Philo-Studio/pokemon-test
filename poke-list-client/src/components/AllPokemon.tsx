import api from '@/server';

const AllPokemon = async () => {
  const pokemon = await api.getAllPokemon();

  return (
    <>
      <h5>All Pokemon</h5>
      <div className="flex flex-wrap gap-4 bg-gray-500 p-4 rounded-md">
        {pokemon.map((p) => (
          <div className="p-2 bg-gray-200 rounded-md" key={p.pokedex_id}>
            <p className="text-sm text-gray-500">{p.pokedex_id}</p>
            <p className="text-md text-gray-500">{p.name}</p>
          </div>
        ))}
      </div>
    </>
  );
};

export default AllPokemon;
