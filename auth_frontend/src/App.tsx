import { useEffect } from 'react'
import * as grpcWeb from 'grpc-web';
import { SignupRequest, User } from './authpb/auth_pb'
import { SignupServiceClient } from './authpb/AuthServiceClientPb'
import './App.css';

function App() {
  let signupServiceClient = new SignupServiceClient('http://localhost:8000')
  const request = new SignupRequest()
  const user = new User()
  user.setUserName("abednarchuk")
  user.setEmail("andrii.bednarchuk@gmail.com")
  user.setPassword("secret")
  request.setUser(user)

  const metadata: grpcWeb.Metadata = {}

  useEffect(() => {
    signupServiceClient.signUp(request, metadata, (err, res) => {
      try {
        console.log(err)
        console.log(res)
      } catch (err) { }
    })
  }, [])

  return (
    <div>
      test
    </div>
  );
}

export default App;
