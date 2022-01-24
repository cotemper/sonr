# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Sonr < Formula
  desc "The Official Sonr CLI tool for building and deploying services on the Sonr network/blockchain."
  homepage "https://sonr.io"
  version "0.1.1"
  license "GPLv3"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/sonr-io/sonr/releases/download/v0.1.1/sonr-highway-0.1.1-darwin-arm64.tar.gz"
      sha256 "67e115b301a18db563d600cc7be4b6470595ef1a5d15a363ccd479dd31799b2c"

      def install
        bin.install "sonr-cli"
        bin.install "sonr-highway"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/sonr-io/sonr/releases/download/v0.1.1/sonr-highway-0.1.1-darwin-amd64.tar.gz"
      sha256 "1e842265918e9fc8fb248b856a028cab8821d11403d6d5f959d3397efd527041"

      def install
        bin.install "sonr-cli"
        bin.install "sonr-highway"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/sonr-io/sonr/releases/download/v0.1.1/sonr-highway-0.1.1-linux-amd64.tar.gz"
      sha256 "fe9107f1ab05a19b9d7072b92bfebf91abdf0bca4ec331a9b504e4eaf321b334"

      def install
        bin.install "sonr-cli"
        bin.install "sonr-highway"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/sonr-io/sonr/releases/download/v0.1.1/sonr-highway-0.1.1-linux-arm64.tar.gz"
      sha256 "7782712e8a749b71b4e425c63da5ebb565f3d1cfba5a477041079556e4bd4fe7"

      def install
        bin.install "sonr-cli"
        bin.install "sonr-highway"
      end
    end
  end

  def caveats; <<~EOS
    Run `brew info sonr` for more information, or run `sonr --help`.
  EOS
  end
end