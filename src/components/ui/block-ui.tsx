import { useState, useCallback, type CSSProperties, useRef } from "react";
import { LoadingSpinner } from "./loading-spinner";

interface BlockUIHook {
  isBlocked: boolean;
  theme: "white" | "black";
  block: (theme?: "white" | "black") => void;
  unblock: () => void;
}

export const useBlockUI = (
  initialTheme: "white" | "black" = "black",
  initialIsBlocked: boolean = false
): BlockUIHook => {
  const [isBlocked, setIsBlocked] = useState(initialIsBlocked);
  const [theme, setTheme] = useState(initialTheme);

  const block = useCallback((theme = "black") => {
    setIsBlocked(true);
    setTheme(theme === "white" ? "white" : "black");
  }, []);

  const unblock = useCallback(() => {
    setIsBlocked(false);
  }, []);

  if (theme !== "white" && theme !== "black") {
    throw new Error("Invalid theme");
  }

  return {
    isBlocked,
    theme,
    block,
    unblock,
  };
};

interface BlockUIProps {
  isBlocked: boolean;
  theme: "white" | "black";
}

export const BlockUI = ({ isBlocked, theme }: BlockUIProps) => {
  if (!isBlocked) {
    return null;
  }

  const overlayStyle: CSSProperties = {
    position: "absolute",
    top: 0,
    left: 0,
    width: "100%", // Changed from right: 0
    height: "100%", // Changed from bottom: 0
    backgroundColor:
      theme === "white" ? "rgba(255, 255, 255, 0.5)" : "rgba(0, 0, 0, 0.5)",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    zIndex: 1000,
  };

  return (
    <div style={overlayStyle}>
      <LoadingSpinner size={32} theme={"white"} />
    </div>
  );
};
