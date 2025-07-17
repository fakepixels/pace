# Homebrew Formula for Pace CLI
# This file should be placed in a homebrew-tap repository
# Usage: brew install fakepixels/tap/pace

class Pace < Formula
  desc "CLI tool for Pace Capital announcements and updates"
  homepage "https://github.com/fakepixels/pace"
  url "https://github.com/fakepixels/pace/archive/v1.0.0.tar.gz"
  sha256 "" # This will be automatically updated by release process
  license "MIT"
  head "https://github.com/fakepixels/pace.git", branch: "main"

  depends_on "go" => :build

  def install
    # Build the binary
    system "go", "build", *std_go_args(ldflags: "-s -w -X main.version=#{version}"), "main.go"
    
    # Rename binary to pace if needed
    if bin/"main".exist?
      mv bin/"main", bin/"pace"
    end
  end

  test do
    # Test that the binary runs and shows version
    assert_match version.to_s, shell_output("#{bin}/pace --version")
    
    # Test help command
    assert_match "Pace CLI", shell_output("#{bin}/pace --help")
  end
end