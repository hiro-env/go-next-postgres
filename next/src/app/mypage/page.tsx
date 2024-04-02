"use client";

import React, { useEffect, useState, useRef } from 'react';
import Header from '../components/Header';
import useAuthCheck from '../hooks/useAuthCheck';
import { useRouter } from 'next/navigation';

const MyPage: React.FC = () => {
  const [nickname, setNickname] = useState('');
  const [profilePic, setProfilePic] = useState('https://via.placeholder.com/150');
  const fileInputRef = useRef<HTMLInputElement>(null);
  const isLoading = useAuthCheck();
  const router = useRouter();

  useEffect(() => {
    if (!isLoading) {
      router.prefetch('/login');
      router.prefetch('/users');
    }
  }, [isLoading, router]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/mypage', {
          credentials: 'include',
        });
        if (!response.ok) throw new Error('failed loading')
        const data = await response.json();
        console.log(data);
  
        setNickname(data.username);
        if (data.image !== "") {
          setProfilePic(`data:image/jpeg;base64,${data.image}`);
        }
      } catch (error) {
        console.error('error', error)
      }
    };
    fetchData();
 }, []);

  const handleNicknameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNickname(e.target.value);
  };

  const handleProfilePicChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = () => {
        setProfilePic(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSave = async () => {

    console.log('Nickname:', nickname);
    console.log('Profile Picture:', profilePic);

    if (profilePic === "https://via.placeholder.com/150") {
      alert("画像をクリックして設定してください。");
      return;
    }

    const base64ImageContent = profilePic.replace(/^data:image\/[a-z]+;base64,/, '');
    const payload = {
      username: nickname,
      image: base64ImageContent
    };
    console.log(payload);

    try {
      const response = await fetch('http://localhost:8080/updateuser', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload),
        credentials: 'include',
      });
      if (response.ok) {
        console.log('データ送信成功');
        alert('マイページを更新しました。')
      } else {
       console.error('データ送信失敗');
       alert('マイページの更新に失敗しました。')
      }
  } catch (error) {
    console.error('送信中にエラーが発生しました', error);
  }
  };

  const handleDelete = async () => {
    const confirm: boolean = window.confirm("本当にアカウントを削除しますか？");
    if (!confirm) return

    const response = await fetch('http://localhost:8080/delete', {
      credentials: 'include',
    });
    console.log(response.status);
    console.log(response.statusText)
    const data = await response.json();
    console.log(data);
    if (response.ok) {
      alert('アカウントを削除しました。')
      router.push('/register')
    } else {
      alert('アカウント削除に失敗しました。')
    }
  }

  return (
    isLoading ? (
      <div>Loading..</div>
    ) : (
    <div className="flex flex-col min-h-screen">
      <Header />
      <div className="flex-grow flex justify-center items-center">
        <div className="bg-white p-8 rounded-lg shadow-md">
          <h1 className="text-3xl font-bold mb-4 text-center">マイページ</h1>
          <div className="flex items-center">
            <div className="relative">
              <label htmlFor="profile-pic-input" className="cursor-pointer">
                <img
                  src={profilePic}
                  alt="プロフィール画像"
                  className="w-32 h-32 rounded-full mr-8"
                />
              </label>
              <input
                id="profile-pic-input"
                type="file"
                accept="image/*"
                ref={fileInputRef}
                onChange={handleProfilePicChange}
                style={{ display: 'none' }}
              />
            </div>
            <div>
              <label htmlFor="nickname" className="block mb-2 font-bold">
                ニックネーム
              </label>
              <input
                type="text"
                id="nickname"
                value={nickname}
                onChange={handleNicknameChange}
                className="w-full px-4 py-2 mb-4 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
              <button
                onClick={handleSave}
                className="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
              >
                保存
              </button>
              <button
                  onClick={handleDelete}
                  className="px-4 py-2 ml-6 bg-red-500 text-white rounded-lg hover:bg-red-600"
                >
                  アカウント削除
                </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    )
  );
};

export default MyPage;