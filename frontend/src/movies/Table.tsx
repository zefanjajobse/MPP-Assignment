import * as React from "react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { MoviesApi } from "../api/MoviesApi";
import "./Table.css";

export function Table() {
  const {
    data,
    error,
    fetchNextPage,
    hasNextPage,
    isFetching,
    isLoading,
    isError,
  } = useInfiniteQuery(
    ["movieList"],
    ({ pageParam = 0 }) => MoviesApi.get({ pageParam }),
    {
      getNextPageParam: (lastPage) => lastPage.offset,
    }
  );

  const movieList = React.useMemo(
    () => (data ? data?.pages.flatMap((item) => item.results) : []),
    [data]
  );

  const observer = React.useRef<IntersectionObserver>();
  const lastElementRef = React.useCallback(
    (node: HTMLDivElement) => {
      if (isLoading) return;
      if (observer.current) observer.current.disconnect();
      observer.current = new IntersectionObserver((entries) => {
        if (entries[0].isIntersecting && hasNextPage && !isFetching) {
          fetchNextPage();
        }
      });
      if (node) observer.current.observe(node);
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [isLoading, hasNextPage]
  );

  return (
    <>
      <h3 className="title">Movies</h3>
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
            {isError && <>{error}</>}
            {!isLoading &&
              !isError &&
              movieList.map((item, i) => (
                <tr
                  key={i}
                  ref={movieList.length === i + 1 ? lastElementRef : null}
                >
                  <td>{item?.imdb_id}</td>
                  <td>{item?.title}</td>
                  <td>{item?.rating}</td>
                  <td>{item?.year}</td>
                </tr>
              ))}
            {isFetching && (
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
