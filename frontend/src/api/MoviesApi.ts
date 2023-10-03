import JsonClient from "./JsonApi";
import { IMovies, IMovieInfo } from "./ReturnTypes";

export class ApiProvider extends JsonClient {
  async get({
    pageParam,
  }: {
    pageParam: number
  }): Promise<IMovies> {
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

  async add(movie: IMovieInfo): Promise<IMovieInfo> {
    return await this.postJsonMethod("movies", movie);
  }
}

export const MoviesApi = new ApiProvider();
