"use client"

import React, { useEffect, useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';

interface userRegisterRequest {
  username: string;
  password: string;
}

const RegisterPage: React.FC = () => {
  const router = useRouter();
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  useEffect(() => {
    const verifyUser = async () => {
      try {
        const response = await fetch('http://localhost:8080/verify', {
          credentials: 'include',
        });
        if (response.ok) {
          router.push('/users');
        }
      } catch (error) {
        console.error(error);
      }
    };
    verifyUser();
  }, []); 

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log('登録情報:', { username, password });
    if (password === "" || username === "") {
      alert("ユーザー名・パスワードを入力してください");
      return;
    }

    const requestData: userRegisterRequest = {
      username,
      password
    }

    const response = await fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(requestData),
      credentials: 'include',
    })

    const data = await response.json();
    console.log(data);
    console.log(response.status);
    if (response.status === 200) {
      router.push('/users');
    } else if (response.status === 409) {
      alert("すでに使用されているユーザーネームです。")
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-lg px-8 py-6 mt-4 text-left bg-white shadow-lg">
        <h3 className="text-2xl font-bold text-center">新規登録</h3>
        <form onSubmit={handleSubmit}>
          <div>
            <label htmlFor="userId" className="block">ユーザーID</label>
            <input
              type="text"
              id="userId"
              placeholder="ユーザーネーム"
              onChange={(e) => setUsername(e.target.value)}
              className="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600"
              value={username}
            />
          </div>
          <div className="mt-4">
            <label htmlFor="password" className="block">パスワード</label>
            <input
              type="password"
              id="password"
              placeholder="パスワード"
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600"
              value={password}
            />
          </div>
          <div className="flex items-center justify-between mt-4">
            <button type="submit" className="px-6 py-2 text-white bg-green-600 rounded-lg hover:bg-green-900">
              新規登録
            </button>
            <Link href="/login" className="px-6 py-2 text-white bg-blue-600 rounded-lg hover:bg-blue-900">
              ログイン画面へ
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
};

export default RegisterPage;