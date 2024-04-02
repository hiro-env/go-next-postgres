"use client"

import React, { useEffect, useState } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { userAgentFromString } from 'next/server';

const Login: React.FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const router = useRouter();

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
    console.log('ログイン情報:', { username, password });

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ username: username, password }),
        credentials: 'include',
      });
  
      const data = await response.json();
  
      if (response.ok) {
        console.log('ログイン成功:', data);
        router.push('/users')
      } else {
        console.error('ログイン失敗:', data.message);
        alert(data.message)
      }
    } catch (error) {
      console.error('エラーが発生しました:', error);
      alert("エラーです。")
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <div className="w-full max-w-lg px-8 py-6 mt-4 text-left bg-white shadow-lg">
        <h3 className="text-2xl font-bold text-center">ログイン</h3>
        <form onSubmit={handleSubmit}>
          <div className="mt-4">
            <div>
              <label className="block" htmlFor="userId">ユーザーID</label>
              <input
                type="text"
                placeholder="ユーザーID"
                id="userId"
                onChange={(e) => setUsername(e.target.value)}
                className="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600"
                value={username}
              />
            </div>
            <div className="mt-4">
              <label className="block" htmlFor="password">パスワード</label>
              <input
                type="password"
                placeholder="パスワード"
                id="password"
                onChange={(e) => setPassword(e.target.value)}
                className="w-full px-4 py-2 mt-2 border rounded-md focus:outline-none focus:ring-1 focus:ring-blue-600"
                value={password}
              />
            </div>
            <div className="flex items-center justify-between mt-4">
              <button className="px-6 py-2 text-white bg-blue-600 rounded-lg hover:bg-blue-900">ログイン</button>
              <Link href="/register" className="px-6 py-2 text-white bg-green-600 rounded-lg hover:bg-green-900">新規登録へ</Link>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;

