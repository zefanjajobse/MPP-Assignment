export interface IMovieInfo {
  imdb_id: string;
  rating: number;
  title: string;
  year: number;
}

export interface IMovies {
  results: IMovieInfo[];
  offset: number | null;
}


export interface ITotalCount {
  total: number;
}
