'use client';

import { useLogin } from '@/hooks/useAuth';
import { useState } from 'react';

export default function LoginForm() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const loginMutation = useLogin();

  const handleLogin = () => {
    loginMutation.mutate({ email, password }, {
      onSuccess: (data) => {
        window.location.replace("/")
      },
      onError: (err) => {
        console.error('Login error:', err);
      }
    });
  };

  return (
    <div>
      <input placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} />
      <input placeholder="Password" type="password" value={password} onChange={e => setPassword(e.target.value)} />
      <button onClick={handleLogin}>Login</button>
    </div>
  );
}