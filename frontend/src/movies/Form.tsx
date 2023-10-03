import * as React from "react";
import { MoviesApi } from "../api/MoviesApi";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export function Form() {
  const queryClient = useQueryClient();
  const [imdb_id, setImdb_id] = React.useState<string>("");
  const [title, setTitle] = React.useState<string>("");
  const [year, setYear] = React.useState<number>(0);
  const [rating, setRating] = React.useState<number>(0);

  const AddMovie = useMutation(
    (v: { imdb_id: string; title: string; rating: number; year: number }) =>
      MoviesApi.add(v),
    {
      // Always refetch after error or success:
      onSettled: () => {
        queryClient.invalidateQueries(["movieList"]);
      },
    }
  );

  const isDisabled =
    imdb_id === "" || title === "" || year === 0 || rating === 0;

  return (
    <div>
      <label>
        imdb_id:
        <input
          value={imdb_id}
          onChange={(e) => setImdb_id(e.target.value)}
          type="text"
        />
      </label>
      <br />
      <label>
        title:
        <input
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          type="text"
        />
      </label>
      <br />
      <label>
        rating:
        <input
          value={rating}
          onChange={(e) => setRating(parseInt(e.target.value))}
          type="text"
        />
      </label>
      <br />
      <label>
        year:
        <input
          value={year}
          onChange={(e) => setYear(parseInt(e.target.value))}
          type="text"
        />
      </label>
      <br />
      <button
        disabled={isDisabled}
        onClick={() =>
          AddMovie.mutate({
            imdb_id,
            title,
            year,
            rating,
          })
        }
      >
        Add
      </button>
    </div>
  );
}
