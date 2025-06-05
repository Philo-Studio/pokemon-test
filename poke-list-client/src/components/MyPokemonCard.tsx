'use client';

import { useState } from 'react';

const MyPokemonCard = ({
  pokemon,
  children,
}: {
  pokemon: { id: string; pokedex_id: string; nickname: string };
  children: React.ReactNode;
}) => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <>
      <div
        className="p-2 bg-gray-200 rounded-md cursor-pointer hover:bg-gray-300 transition-colors"
        key={pokemon.id}
        onClick={() => setIsOpen(true)}
      >
        <p className="text-sm text-gray-500">{pokemon.pokedex_id}</p>
        <p className="text-md text-gray-500">{pokemon.nickname}</p>
      </div>
      {isOpen && (
        <div className="z-10 fixed h-[90vh] w-[90vw] border-2 border-gray-500 bg-gray-900 rounded-md shadow-2xl top-[5vh] left-[5vw] p-8">
          <div className="z-10">
            <div className="flex justify-between">
              <h1 className="text-2xl font-bold">Details</h1>
              <button onClick={() => setIsOpen(false)} className="bg-gray-500 p-2 rounded-md">
                Close
              </button>
            </div>
            {children}
          </div>
        </div>
      )}
    </>
  );
};

export default MyPokemonCard;
