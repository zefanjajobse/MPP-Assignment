import JsonClient from "./JsonApi";

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

export class ApiProvider extends JsonClient {
  async get({
    pageParam,
  }: {
    pageParam: number
  }): Promise<any> {
    const limit = 400;
    const offset = pageParam ? pageParam - 1 : 0;
    const data = await this.getJsonMethod("movies", {
      Offset: offset,
      Limit: limit,
    });

    return {
      results: data,
      offset: offset + limit,
    };
  }
}

export const MoviesApi = new ApiProvider();
