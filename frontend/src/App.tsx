import * as React from "react";
import { useInfiniteQuery } from "@tanstack/react-query";
import { MoviesApi } from "./api/MoviesApi";

import "./App.css";

function App() {
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

  // MoviesApi.add({
  //   imdb_id: "tt0111161",
  //   title: "The Shawshank Redemption",
  //   rating: 9.2,
  //   year: 1994,
  // });

  return (
    <div style={{ maxHeight: "400px", overflowY: "auto", marginTop: "8px" }}>
      <table style={{ borderCollapse: "collapse", width: "100%" }}>
        <thead
          style={{
            position: "sticky",
            top: "0",
            backgroundColor: "white",
            textAlign: "left",
          }}
        >
          <th>imdb_id</th>
          <th>title</th>
          <th>rating</th>
          <th>year</th>
        </thead>
        <tbody>
          {isError && <>{error}</>}
          {!isLoading &&
            !isError &&
            movieList.map((item, i) => (
              <tr ref={movieList.length === i + 1 ? lastElementRef : null}>
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
  );
}

export default App;
