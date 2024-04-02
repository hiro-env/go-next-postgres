"use client"

import React, { useEffect } from 'react';
import { useState } from 'react';
import Search from '../components/Search';
import Header from '../components/Header'
import useAuthCheck from '../hooks/useAuthCheck';
import { useRouter } from 'next/navigation';

interface User {
  id: number;
  nickname: string;
  profilePic: string;
}

interface UserResponse {
  id: number;
  nickname: string;
  image: string;
}

interface UserSearchRequest {
  searchTerm: string;
}

const UserList: React.FC = () => {
  const isLoading = useAuthCheck();
  const [filteredUsers, setFilteredUsers] = useState<User[]>([]);
  const router = useRouter();

  useEffect(() => {
    if (!isLoading) {
      router.prefetch('/login');
      router.prefetch('/users');
    }
  }, [isLoading, router]);

  useEffect(() => {
    const fetchUsers = async () => {
      const response = await fetch('http://localhost:8080/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          searchTerm: "",
        }),
        credentials: 'include',
      })
      if (!response.ok) {
        return;
      }
      const data: UserResponse[] = await response.json();
      const users: User[] = data.map((UserResponse) => {
        UserResponse.image = (UserResponse.image !== null && UserResponse.image !== "") ? 
        `data:image/jpeg;base64,${UserResponse.image}` : 'https://via.placeholder.com/150'
        
        return {
          id: UserResponse.id,
          nickname: UserResponse.nickname,
          profilePic: UserResponse.image,
        }
      });
      console.log(users);
      setFilteredUsers(users);
    };
    fetchUsers();
  }, []);

  const handleSearch = (searchTerm: string) => {
    searchUsers(searchTerm);

    setFilteredUsers(
      filteredUsers.filter((user: User) =>
        user.nickname.toLowerCase().includes(searchTerm.toLowerCase())
      )
    );
  };

  const searchUsers = async (searchTerm: string) => {
    const requestData: UserSearchRequest = {
      searchTerm: searchTerm
    }

    const response = await fetch('http://localhost:8080/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestData),
      credentials: 'include',
    })
    if (response.status === 401) {
      window.location.href = '/login'
      return;
    }

    const data: UserResponse[] = await response.json();
    console.log(data);

    const users: User[] = data.map((UserResponse) => {
      UserResponse.image = (UserResponse.image !== null && UserResponse.image !== "") ? 
      `data:image/jpeg;base64,${UserResponse.image}` : 'https://via.placeholder.com/150'
      
      return {
        id: UserResponse.id,
        nickname: UserResponse.nickname,
        profilePic: UserResponse.image,
      }
    });
    console.log(users);
    setFilteredUsers(users);
  }

  return (
    isLoading ? (
      <div>Loading...</div>
    ) : (
      <div className="flex flex-col min-h-screen">
      <Header />
      <div className="sticky top-0 bg-white z-10">
        <div className="flex flex-col items-center mt-8">
          <Search onSearch={handleSearch} />
          <h1 className="text-3xl font-bold mb-4">ユーザー一覧</h1>
        </div>
      </div>
      <div className="flex-grow overflow-auto">
        <div className="flex flex-col items-center">
          <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-4xl">
            <ul>
              {filteredUsers.map((user) => (
                <li key={user.id} className="flex items-center justify-between mb-4">
                  <div className="flex items-center">
                    <img src={user.profilePic} alt="画像" className="w-10 h-10 rounded-full mr-4" />
                    <span>{user.nickname}</span>
                  </div>
                </li>
              ))}
            </ul>
          </div>
        </div>
      </div>
    </div>
    )
  );
};

export default UserList;