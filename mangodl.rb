# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!
class Mangodl < Formula
  desc "An easy-to-use cli tool for downloading manga"
  homepage "https://github.com/gyro7/mangodl"
  url "https://github.com/gyro7/mangodl/releases/download/mangodl-v1.3-mac/mangodl.tar.gz"
  sha256 "5213bd5b4a3faa2246ec98ad0a3e772989bba9957294ffd5c81b2b48b87d243f"
  license "GPL-3.0-only"

  depends_on "go" => :build

  def install
      system "go", "build", "-o", "#{prefix}"
  end

  test do
      system "mangodl", "--help"
  end
end
