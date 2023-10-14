import * as React from "react";
import { UseQueryResult, useQuery } from "@tanstack/react-query";
import { MoviesApi } from "../api/MoviesApi";
import "./Table.css";
import { IMovies } from "../api/ReturnTypes";

export function Table() {
  const limit = 20;
  const [page, setPage] = React.useState(1);

  const { data: totalCount } = useQuery(["movieCount"], () =>
    MoviesApi.count()
  );
  const pageCount = React.useMemo(
    () => Math.ceil(totalCount?.total / limit) | 0,
    [totalCount]
  );

  const { data, error, isError, isLoading }: UseQueryResult<IMovies, any> =
    useQuery(["movieList", { page }], () =>
      MoviesApi.get({ offset: (page - 1) * limit, limit })
    );

  return (
    <>
      <h3 className="title">Movies</h3>
      <div className="center">
        <button onClick={() => setPage(1)} disabled={page === 1}>
          First
        </button>
        <button
          onClick={() => setPage((prev) => Math.max(prev - 1, 1))}
          disabled={page === 1}
        >
          Prev
        </button>
        <div className="pageSelector">
          {page > 3 && <button disabled>...</button>}
          {page > 2 && (
            <button onClick={() => setPage(page - 2)}>{page - 2}</button>
          )}
          {page > 1 && (
            <button onClick={() => setPage(page - 1)}>{page - 1}</button>
          )}
          <button disabled>{page}</button>
          {page < pageCount && (
            <button onClick={() => setPage(page + 1)}>{page + 1}</button>
          )}
          {page < pageCount - 1 && (
            <button onClick={() => setPage(page + 2)}>{page + 2}</button>
          )}
          {page < pageCount - 2 && <button disabled>...</button>}
        </div>
        <button
          onClick={() => setPage((prev) => Math.min(prev + 1, pageCount))}
          disabled={page === pageCount}
        >
          Next
        </button>
        <button
          onClick={() => setPage(pageCount)}
          disabled={page === pageCount}
        >
          Last
        </button>
      </div>
      <div className="table">
        <table>
          <thead>
            <tr>
              <th>imdb_id</th>
              <th>title</th>
              <th>rating</th>
              <th>year</th>
            </tr>
          </thead>
          <tbody>
            {isError && (
              <tr>
                <td>Something went wrong...</td>
                <td>{error?.error?.message?.message}</td>
                <th>N/A</th>
                <th>N/A</th>
              </tr>
            )}
            {data?.results?.map((item, i) => (
              <tr key={i}>
                <td>{item?.imdb_id}</td>
                <td>{item?.title}</td>
                <td>{item?.rating}</td>
                <td>{item?.year}</td>
              </tr>
            ))}
            {isLoading && (
              <tr>
                <td>loading...</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
}
