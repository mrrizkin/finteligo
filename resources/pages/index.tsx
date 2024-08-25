import { Link } from "react-router-dom";

export default function IndexPage() {
  return (
    <div>
      <h1>Welcome to Vite</h1>
      <p>
        Edit <code>resources/pages/index.tsx</code> and save to test HMR updates.
      </p>
      <Link to="/dashboard">Go to dashboard</Link>
    </div>
  );
}
