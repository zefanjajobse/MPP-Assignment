import JsonClient from "./JsonApi";
import { IMovies, IMovieInfo, ITotalCount } from "./ReturnTypes";

export class ApiProvider extends JsonClient {
  async get({
    offset,
    limit,
  }: {
    offset: number
    limit: number
  }): Promise<IMovies> {
    const data = await this.getJsonMethod("movies", {
      Offset: offset,
      Limit: limit,
    });

    return {
      results: data,
      offset: offset + limit,
    };
  }

  async count(): Promise<ITotalCount> {
    return await this.getJsonMethod("movies/count", {});
  }

  async add(movie: IMovieInfo): Promise<IMovieInfo> {
    return await this.postJsonMethod("movies", movie);
  }
}

export const MoviesApi = new ApiProvider();
