DROP TABLE IF EXISTS my_pokemon;

CREATE TABLE my_pokemon
(
  id                          UUID      DEFAULT gen_random_uuid() NOT NULL
    CONSTRAINT my_pokemon_pk
      PRIMARY KEY,
  pokedex_id                  INT                                 NOT NULL,
  nickname                    TEXT                                NOT NULL,
  created_at                  TIMESTAMP DEFAULT NOW()             NOT NULL,
  deleted_at                  TIMESTAMP
)
;


INSERT INTO public.my_pokemon ( pokedex_id, nickname )
VALUES ( 50, 'Tony' ),
       ( 2, 'Carl' ),
       ( 60, 'Mike' )
;
