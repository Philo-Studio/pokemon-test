import AllPokemon from '@/components/AllPokemon';
import MyPokemon from '@/components/MyPokemon';
import { Suspense } from 'react';

export default function Home() {
  return (
    <div className="flex flex-col gap-4 max-w-[80vw] mx-auto">
      <Suspense
        fallback={
          <>
            <h5>My Pokemon</h5>
            <div className="flex flex-wrap gap-4 bg-gray-500 p-4 rounded-md">
              <h6 className="text-center">Loading...</h6>
            </div>
          </>
        }
      >
        <MyPokemon />
      </Suspense>
      <Suspense
        fallback={
          <>
            <h5>All Pokemon</h5>
            <div className="flex flex-wrap gap-4 bg-gray-500 p-4 rounded-md">
              <h6 className="text-center">Loading...</h6>
            </div>
          </>
        }
      >
        <AllPokemon />
      </Suspense>
    </div>
  );
}
