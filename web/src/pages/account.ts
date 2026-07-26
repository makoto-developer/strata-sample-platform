import { gql, request } from '../lib/gql.js';

// マイページ: ユーザーと注文履歴(federation でサブグラフを横断する)
const ACCOUNT_QUERY = gql`
  query Account($id: ID!) {
    user(id: $id) {
      id
      name
      orders {
        id
        status
      }
    }
  }
`;

const UPDATE_PROFILE = gql`
  mutation UpdateProfile($id: ID!, $name: String!) {
    updateProfile(id: $id, name: $name) {
      id
      name
    }
  }
`;

export async function loadAccount(id: string): Promise<unknown> {
  return request(ACCOUNT_QUERY, { id });
}

export async function saveProfile(id: string, name: string): Promise<unknown> {
  return request(UPDATE_PROFILE, { id, name });
}
